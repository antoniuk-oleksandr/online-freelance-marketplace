<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";

    const {children}: LayoutProps = $props();
</script>

<div class="flex w-full h-svh items-center">
    <PaperElement
            styles="w-full shadow-none flex h-full md:h-form rounded-none lg:rounded-lg !p-0 grid grid-cols-1 lg:grid-cols-2 overflow-hidden !ring-0 lg:!ring-1 !bg-transparent lg:!bg-light-palette-background-block lg:dark:!bg-dark-palette-background-block">
        {@render children()}
    </PaperElement>
</div>