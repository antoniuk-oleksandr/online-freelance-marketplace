<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="p-6 w-full flex flex-col gap-y-3 bg-opacity-90 text-start absolute left-0 bottom-0 bg-light-palette-background-block dark:bg-dark-palette-background-block">
    {@render children()}
</div>