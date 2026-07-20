// Copyright (c) Microsoft Corporation.
// Copyright (c) 2026 Josh Mouch.
// Licensed under the MIT License.

package agency.open.ahp;

/**
 * Entry point that replays the AHP reducer corpus against the extracted core.
 *
 * <p>The reducers reached from here are not hand-written: they are mechanically
 * extracted from a Dafny specification whose convergence and fold properties are
 * machine-checked. Running this class replays the fixture corpus for all seven
 * AHP channels and prints a per-channel pass count, which is the cheapest way to
 * confirm that a build of this artifact behaves like the verified specification.
 *
 * <p>Equivalent to the {@code Main} method of the Dafny {@code ClientMain} module.
 */
public final class CorpusRunner {

    private CorpusRunner() {
        // Not instantiable.
    }

    public static void main(String[] args) {
        dafny.Helpers.withHaltHandling(
                () ->
                        agency.open.ahp.ClientMain.__default.__Main(
                                dafny.Helpers.UnicodeFromMainArguments(args)));
    }
}
