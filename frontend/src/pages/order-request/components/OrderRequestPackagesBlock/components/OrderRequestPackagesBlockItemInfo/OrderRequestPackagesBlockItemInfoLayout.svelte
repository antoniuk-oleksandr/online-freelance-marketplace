<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="flex items-center gap-1 text-light-palette-text-secondary dark:text-dark-palette-text-secondary text-sm">
    {@render children()}
</div>