<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";

    const {children}: LayoutProps = $props();
</script>

<div class="flex h-svh items-center">
    <PaperElement
            styles="w-full flex h-form rounded-none lg:rounded-lg !bg-light-palette-background-block dark:!bg-dark-palette-background-block !p-0 grid grid-cols-1 lg:grid-cols-2 overflow-hidden border-none lg:border-solid bg-transparent">
        {@render children()}
    </PaperElement>
</div>