/**
 * Dafny Generator — emits Dafny `datatype`s for the AHP wire types from the
 * TypeScript source of truth (parsed via ts-morph), sibling of generate-rust.ts.
 * Per the plan: emits the TYPES; reducers are hand-authored + proven.
 * Run from the fork (for ts-morph): cd <fork> && node_modules/.bin/tsx <this> --out <path>
 */
import { Project, InterfaceDeclaration, EnumDeclaration, PropertySignature } from 'ts-morph';
import path from 'path';
import fs from 'fs';

const FORK = process.env.AHP_REPO;
if (!FORK) throw new Error('AHP_REPO must be set to an agent-host-protocol checkout (no hard-coded default)');
const TYPES_DIR = path.join(FORK, 'types');

function stripIPrefix(n: string): string { return n.replace(/^I(?=[A-Z])/, ''); }

function mapType(tsType: string): string {
  tsType = tsType.replace(/import\([^)]+\)\./g, '').trim();
  while (tsType.startsWith('(') && tsType.endsWith(')')) tsType = tsType.slice(1, -1).trim();
  if (tsType === 'string') return 'string';
  if (tsType === 'number') return 'int';
  if (tsType === 'boolean') return 'bool';
  if (tsType === 'unknown' || tsType === 'object' || tsType === 'JsonPrimitive') return 'Json';
  if (tsType === 'true' || tsType === 'false') return 'bool';
  if (tsType === 'URI') return 'string';
  if (tsType === 'SessionStatus') return 'SessionStatus';
  const nullMatch = tsType.match(/^(.+?)\s*\|\s*null$/);
  if (nullMatch) return `Option<${mapType(nullMatch[1])}>`;
  const undefMatch = tsType.match(/^(.+?)\s*\|\s*undefined$/);
  if (undefMatch) return `Option<${mapType(undefMatch[1])}>`;
  const arr = tsType.match(/^(.+)\[\]$/); if (arr) return `seq<${mapType(arr[1])}>`;
  const arrG = tsType.match(/^Array<(.+)>$/); if (arrG) return `seq<${mapType(arrG[1])}>`;
  const rec = tsType.match(/^Record<string,\s*(.+)>$/);
  if (rec) { const inner = rec[1].trim(); return (inner === 'unknown' || inner === 'never') ? 'map<string, Json>' : `map<string, ${mapType(inner)}>`; }
  const enumMember = tsType.match(/^(\w+)\.(\w+)$/); if (enumMember) return stripIPrefix(enumMember[1]);
  const enumUnion = tsType.match(/^(\w+)\.\w+(\s*\|\s*\1\.\w+)*$/); if (enumUnion) return stripIPrefix(enumUnion[1]);
  if (tsType.startsWith("'") && tsType.endsWith("'")) return 'string';
  if (/^'[^']*'(\s*\|\s*'[^']*')+$/.test(tsType)) return 'string';
  if (tsType.startsWith('{')) return 'Json';
  if (/[|<>&]/.test(tsType)) return 'Json';
  return stripIPrefix(tsType);
}

const RESERVED = new Set([
  // types + collections
  'type','nat','int','real','bool','char','string','object','set','iset','multiset','map','imap','seq','array','array2','array3','bv','ORDINAL',
  // declarations
  'const','var','method','function','predicate','lemma','ghost','class','trait','datatype','codatatype','newtype','module','import','export','opened','include','abstract','iterator',
  // spec clauses
  'requires','ensures','modifies','reads','decreases','invariant','yields','yield','provides','reveals','witness',
  // statements / expressions
  'match','case','assert','assume','expect','print','return','break','continue','new','this','old','fresh','forall','exists','if','then','else','while','in','calc','by','label','unchanged','allocated',
  // literals
  'true','false','null',
  // built-ins we emit
  'value','values','fields','elems',
]);
function fieldName(n: string): string { return RESERVED.has(n) ? `${n}_` : n; }

function isDiscriminant(p: PropertySignature): boolean {
  if (p.getName() !== 'type') return false;
  const t = p.getTypeNode()?.getText() ?? '';
  return t.startsWith("'") || t.includes('.');
}

function emitInterface(iface: InterfaceDeclaration): string {
  const name = stripIPrefix(iface.getName());
  // Keep only plain-identifier properties; drop discriminants + TS computed/mapped keys ([X.Y]:…).
  const props = iface.getProperties().filter(p => !isDiscriminant(p) && /^[A-Za-z_]\w*$/.test(p.getName()));
  if (props.length === 0) return `  datatype ${name} = ${name}(placeholder_: bool)`;  // Dafny needs ≥1 field
  const seen = new Set<string>();
  const fields = props.map(p => {
    const t = p.getTypeNode()?.getText() ?? 'unknown';   // syntactic only (no semantic resolution)
    const dt = p.hasQuestionToken() && !t.includes('undefined') ? `Option<${mapType(t)}>` : mapType(t);
    let fn = fieldName(p.getName());
    while (seen.has(fn)) fn = fn + '_';
    seen.add(fn);
    return `${fn}: ${dt}`;
  });
  return `  datatype ${name} = ${name}(${fields.join(', ')})`;
}

function emitEnum(e: EnumDeclaration, typeNames: Set<string>, ctorNames: Set<string>): string {
  const name = stripIPrefix(e.getName());
  const variants = e.getMembers().map(m => {
    let v = m.getName().charAt(0).toUpperCase() + m.getName().slice(1);
    // Suffix any variant that collides with a TYPE name or an already-used constructor
    // (Dafny constructors share one namespace; a variant can't shadow a datatype or repeat).
    while (typeNames.has(v) || ctorNames.has(v)) v = v + '_';
    ctorNames.add(v);
    return v;
  });
  const uniq = Array.from(new Set(variants));
  return `  datatype ${name} = ${uniq.join(' | ')}`;
}

// Syntactic-only: no tsconfig / no @types / no lib walk-up (avoids permission walk to /Users/josh/node_modules).
const project = new Project({ skipLoadingLibFiles: true, compilerOptions: { types: [], typeRoots: [], noLib: true, skipLibCheck: true } });
project.addSourceFilesAtPaths(path.join(TYPES_DIR, '**/*.ts'));
const header = [
  '// GENERATED by gen/generate-dafny.ts from the AHP types/ — DO NOT EDIT.',
  '// Wire-type datatypes only; reducers are hand-authored + proven.',
  'module AhpMirror {',
  '  datatype Option<T> = None | Some(value: T)',
  '  datatype Json = JNull | JBool(b: bool) | JNum(n: int) | JStr(s: string) | JArr(elems: seq<Json>) | JObj(fields: map<string, Json>)',
  '  type SessionStatus = bv32',
  '',
];
// PASS 1 — collect every type name (so enum variants can avoid shadowing a datatype).
const typeNames = new Set<string>(['Option', 'Json', 'SessionStatus']);
const sfList = project.getSourceFiles().filter(sf => sf.getFilePath().includes('/types/') && !sf.getFilePath().includes('.test.'));
for (const sf of sfList) {
  for (const e of sf.getEnums()) typeNames.add(stripIPrefix(e.getName()));
  for (const iface of sf.getInterfaces()) if (iface.isExported()) typeNames.add(stripIPrefix(iface.getName()));
  for (const ta of sf.getTypeAliases()) if (ta.isExported()) typeNames.add(stripIPrefix(ta.getName()));
}
// Constructor namespace: every datatype's constructor == its own name, plus the built-ins.
const ctorNames = new Set<string>(['None', 'Some', 'JNull', 'JBool', 'JNum', 'JStr', 'JArr', 'JObj', ...typeNames]);

// PASS 2 — emit.
const body: string[] = [];
const emittedNames = new Set<string>(['Option', 'Json', 'SessionStatus']);
let nIface = 0, nEnum = 0, nAlias = 0, nAliasUnion = 0, nNever = 0;
for (const sf of sfList) {
  for (const e of sf.getEnums()) {
    const nm = stripIPrefix(e.getName());
    if (emittedNames.has(nm)) continue; emittedNames.add(nm);
    body.push(emitEnum(e, typeNames, ctorNames)); nEnum++;
  }
  for (const iface of sf.getInterfaces()) {
    if (!iface.isExported()) continue;
    const nm = stripIPrefix(iface.getName());
    if (emittedNames.has(nm)) continue; emittedNames.add(nm);
    body.push(emitInterface(iface)); nIface++;
  }
  // Type aliases. A union whose members are ALL declared types (interfaces /
  // enums / other datatypes) is emitted as a real Dafny DISCRIMINATED UNION,
  // wrapping each member in a uniquely-named variant. Unions that mix in
  // primitives / string-literals / `never` / inline shapes stay opaque stubs
  // (carry the raw Json) — refining those needs primitive-arm modelling.
  for (const ta of sf.getTypeAliases()) {
    if (!ta.isExported()) continue;
    const nm = stripIPrefix(ta.getName());
    if (emittedNames.has(nm)) continue; emittedNames.add(nm);
    const rawTxt = ta.getTypeNode()?.getText()?.replace(/\n/g, ' ').trim() ?? '';
    // split on `|`; drop the empty arm a leading-pipe union (`| A | B`) produces.
    const members = rawTxt.includes('|')
      ? rawTxt.split('|').map(s => stripIPrefix(s.trim())).filter(s => s.length > 0)
      : [];
    const allDeclared = members.length >= 2 &&
      members.every(m => /^[A-Za-z_]\w*$/.test(m) && typeNames.has(m) && m !== nm);
    if (rawTxt === 'never') {
      // TypeScript `never` is uninhabited.  Mapping it to the old opaque
      // `...(raw: Json)` stub silently invented wire values (notably for
      // ServerCanvasAction), which could then be mistaken for server-only
      // protocol variants.  A possibly-empty Dafny subset type preserves the
      // exact trust boundary: it has no constructible values and introduces no
      // semantic action arm.  `witness *` is Dafny's standard declaration form
      // for a possibly-empty subset type; no axiom or assumed inhabitant is
      // introduced.
      body.push(`  type ${nm} = raw: Json | false witness *  // exact TS never: intentionally uninhabited`);
      nNever++;
    } else if (allDeclared) {
      // uniquely-named variant per member; field name derived from the member.
      const seen = new Set<string>();
      const arms = members.map(m => {
        let fld = m.charAt(0).toLowerCase() + m.slice(1);
        if (RESERVED.has(fld)) fld = fld + '_';
        while (seen.has(fld)) fld = fld + '_'; seen.add(fld);
        return `    | ${nm}_${m}(${fld}: ${m})`;
      });
      body.push(`  datatype ${nm} =\n${arms.join('\n')}`);
      nAliasUnion++;
    } else {
      body.push(`  datatype ${nm} = ${nm}(raw: Json)  // opaque: was TS alias \`${rawTxt.slice(0, 60)}\` (non-declared-member union)`);
      nAlias++;
    }
  }
}
const out = process.argv.includes('--out') ? process.argv[process.argv.indexOf('--out') + 1] : path.join(FORK, 'clients/dafny/spec/mirror.generated.dfy');
const content = [...header, ...body, '}', ''].join('\n');

// --check = drift gate (Phase 8): fail if the committed file differs from a fresh regen.
if (process.argv.includes('--check')) {
  const existing = fs.existsSync(out) ? fs.readFileSync(out, 'utf8') : '';
  if (existing !== content) {
    console.error(`DRIFT: ${out} is stale — regenerate. (${nIface}+${nEnum}+${nAliasUnion}+${nAlias}+${nNever} types)`);
    process.exit(1);
  }
  console.log(`fresh: ${out} matches types/ (${nIface} datatypes + ${nEnum} enums + ${nAliasUnion} alias-unions + ${nAlias} alias-stubs + ${nNever} never-types)`);
} else {
  fs.mkdirSync(path.dirname(out), { recursive: true });
  fs.writeFileSync(out, content);
  console.log(`generated ${nIface} datatypes + ${nEnum} enums + ${nAliasUnion} alias-unions + ${nAlias} alias-stubs + ${nNever} never-types → ${out}`);
}
