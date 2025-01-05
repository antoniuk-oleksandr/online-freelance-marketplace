<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";

    const {children}: LayoutProps = $props();
</script>

<PaperElement styles="flex flex-col gap-3">
    {@render children()}
</PaperElement>