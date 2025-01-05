<script lang="ts">
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<PaperElement
        styles="flex text-base flex-col items-center gap-3 p-6 max-w-168 !border-0 md:!border !bg-transparent md:!bg-light-palette-background-block md:dark:!bg-dark-palette-background-block"
>
    {@render children()}
</PaperElement>