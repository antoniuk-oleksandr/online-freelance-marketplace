<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex justify-center gap-x-1 text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    {@render children()}
</div>