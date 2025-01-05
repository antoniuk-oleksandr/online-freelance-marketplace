<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {handleKeepsSignedInClick} from "@/pages/sign-in/handlers.ts";

    type KeepMeSignedInBlockLayoutProps = LayoutProps & {
        keepSignedIn: boolean;
        setKeepSignedIn: (value: boolean) => void;
    }

    const {children, keepSignedIn, setKeepSignedIn}: KeepMeSignedInBlockLayoutProps = $props();
</script>

<button
        type="button"
        onclick={() => handleKeepsSignedInClick(keepSignedIn, setKeepSignedIn)}
        class="flex gap-x-3 items-center">
    {@render children()}
</button>