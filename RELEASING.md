# Releasing

Artifacts for every shipping target are already built, tested, and attached to
[GitHub Releases](https://github.com/joshmouch/ahp-verified/releases). Pushing them to the
language registries is the one step that needs a credential, and each registry needs it once.

Every publish workflow is **manual and confirm-gated** — it runs the full validation (build with
warnings-as-errors, smoke test, pack) and only uploads when you explicitly type `publish`.
Nothing auto-publishes on a tag.

## NuGet — `Ahp.Core.Verified`

1. [nuget.org](https://www.nuget.org) → sign in → avatar → **API Keys** → **Create**
   - Key name: `ahp-verified`
   - Scopes: **Push** → *Push new packages and package versions*
   - Glob pattern: `Ahp.*`
   - Expiration: 365 days (calendar a rotation)
2. Repo → **Settings → Environments → `nuget-org` → Add secret**
   - Name: `NUGET_API_KEY`, value: the key from step 1
3. **Actions → Publish .NET (NuGet) → Run workflow** → type `publish`

Because `Ahp.Core.Verified` does not exist yet, the glob-scoped key is what lets the first push
create it. Consider a [NuGet ID prefix reservation](https://learn.microsoft.com/nuget/nuget-org/id-prefix-reservation)
for `Ahp.*` afterwards.

## npm — `@open-agency/ahp`

1. Create the free org/scope **`open-agency`** on npmjs.com (scoped public packages are free).
2. npmjs.com → avatar → **Access Tokens** → **Generate** → *Automation* token.
3. Repo → **Settings → Environments → `npm` → Add secret** → `NPM_TOKEN`.
4. **Actions → Publish npm → Run workflow** → type `publish`.

The workflow publishes with `--access public --provenance`. It never rebuilds `dist/ahp.cjs` —
regenerating requires Dafny plus a private runtime dependency, so the committed bundle is what
ships.

## PyPI — `agent-host-protocol`

Uses **Trusted Publishing** (OIDC), so there is no long-lived token to store.

1. [pypi.org](https://pypi.org) → **Your projects → Publishing → Add a new pending publisher**
   - PyPI project name: `agent-host-protocol`
   - Owner: `joshmouch` · Repository: `ahp-verified`
   - Workflow: `publish-pypi.yml` · Environment: `pypi`
2. **Actions → Publish PyPI → Run workflow** → type `publish`.

## Maven Central — `agency.open.ahp:ahp-core`

The only registry requiring domain verification.

1. Create a [Central Portal](https://central.sonatype.com) account.
2. Verify the `agency.open` namespace by adding the Portal-generated **DNS TXT record on
   `open.agency`** (you own the domain). Verifying the apex grants `agency.open.*`.
   Zero-domain fallback: use `io.github.joshmouch`, verified via a Portal-named GitHub repo.
3. Generate a GPG key, publish it to a keyserver, and add the key + Portal credentials to
   `~/.m2/settings.xml`.
4. `cd java && mvn -B deploy` (bundle needs jar + `-sources` + `-javadoc` + full POM).

Maven has no automated workflow here yet, deliberately — the signing/verification setup is
interactive and one-time.

## Go — already released

Go resolves modules straight from the git tag; there is no registry push and no credential.
`github.com/joshmouch/ahp-verified/go` is live at tag `go/v0.1.0`:

```sh
go get github.com/joshmouch/ahp-verified/go@v0.1.0
```

## Cutting a new version

1. Bump the version in the target's manifest (`cs/.../*.csproj`, `js/package.json`,
   `py/pyproject.toml`, `java/pom.xml`).
2. Regenerate from the Dafny core if the protocol changed. Regeneration requires
   Dafny 4.11.0 plus the Conflux runtime package, which is not published — see
   [REPRODUCIBILITY.md](REPRODUCIBILITY.md). The gates live in [gen/](gen/); the
   proof sources live in [spec/](spec/).
3. Commit, tag `<lang>/vX.Y.Z`, push. The tag cuts a GitHub Release.
4. Run the corresponding publish workflow with `publish`.
