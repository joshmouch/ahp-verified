"""Agent Host Protocol (AHP) - formally verified reducer core.

This package embeds the AHP channel reducers extracted from a Dafny
specification. The reducers, their convergence properties and the fold
laws over them are machine-checked proofs, not hand-written code.

The extracted modules are held in the private ``_core`` subpackage. They
are generated with flat, absolute module names (``Chat``, ``Session``,
...), so ``_core`` is placed on ``sys.path`` at import time and the
modules are re-exported here under a curated namespace.

Channels
--------
root            Agent / session / terminal registry and client config.
session         Session lifecycle.
chat            Turn lifecycle and the tool-call state machine.
terminal        Terminal streams.
canvas          Canvas surface.
changeset       File changesets.
annotations     Source annotations.
resource_watch  Resource watching.

Each channel module exposes an ``ApplyTo<Channel>(state, action, now)``
reducer returning a ``(next_state, outcome)`` pair, alongside the
datatype constructors for its states and actions.

Example
-------
>>> from agent_host_protocol import root
>>> state = root.RootState_RootState(
...     agents, active_sessions, terminals, config)
>>> next_state, outcome = root.ApplyToRoot(state, action, now)
"""

from __future__ import annotations

import os as _os
import sys as _sys

_CORE_DIR = _os.path.join(_os.path.dirname(_os.path.abspath(__file__)), "_core")

# The extracted modules import one another by bare name; make them resolvable.
if _CORE_DIR not in _sys.path:
    _sys.path.insert(0, _CORE_DIR)

import AhpSkeleton as _AhpSkeleton  # noqa: E402
import Annotations as _Annotations  # noqa: E402
import Canvas as _Canvas  # noqa: E402
import Changeset as _Changeset  # noqa: E402
import Chat as _Chat  # noqa: E402
import ResourceWatch as _ResourceWatch  # noqa: E402
import Session as _Session  # noqa: E402
import Terminal as _Terminal  # noqa: E402
import _dafny  # noqa: E402

#: Root channel - agent/session/terminal registry and client config.
root = _AhpSkeleton
#: Session channel - session lifecycle.
session = _Session
#: Chat channel - turn lifecycle and tool-call state machine.
chat = _Chat
#: Terminal channel - terminal streams.
terminal = _Terminal
#: Canvas channel - canvas surface.
canvas = _Canvas
#: Changeset channel - file changesets.
changeset = _Changeset
#: Annotations channel - source annotations.
annotations = _Annotations
#: ResourceWatch channel - resource watching.
resource_watch = _ResourceWatch

#: The Dafny runtime support module (sequences, maps, bignums, ...).
dafny = _dafny

__version__ = "0.1.0"

CHANNELS = (
    "root",
    "session",
    "chat",
    "terminal",
    "canvas",
    "changeset",
    "annotations",
    "resource_watch",
)

__all__ = ("CHANNELS", "dafny", "__version__", *CHANNELS)


def run_corpus() -> dict[str, tuple[int, int]]:
    """Run the embedded reducer fixture corpus against the extracted code.

    Returns a mapping of channel name to a ``(passed, total)`` pair. This
    is the same corpus the Dafny build checks the extraction against, and
    is useful as an installation self-test.
    """
    results: dict[str, tuple[int, int]] = {}
    for name in CHANNELS:
        module = globals()[name]
        runner = getattr(module.default__, "RunCorpus", None)
        if runner is None:
            continue
        passed, total = runner()
        results[name] = (int(passed), int(total))
    return results
