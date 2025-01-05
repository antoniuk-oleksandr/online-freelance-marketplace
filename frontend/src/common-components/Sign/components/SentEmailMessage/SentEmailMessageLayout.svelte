<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";

    const {children}: LayoutProps = $props();
</script>

<div class="grid h-svh place-items-center">
    <PaperElement
            styles="flex text-base flex-col items-center gap-3 p-6 max-w-168 !border-0 md:!border !bg-transparent md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
    >
        {@render children()}
    </PaperElement>
</div>