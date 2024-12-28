<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";

    const {children}: LayoutProps = $props();
</script>

<PaperElement
        styles={"row-start-2 !p-0 col-start-1 md:row-start-1 md:col-start-2 border border-light-palette-divider dark:border-dark-palette-divider h-fit md:top-22 md:sticky"}
>
    {@render children()}
</PaperElement>