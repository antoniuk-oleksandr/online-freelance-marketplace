<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="p-6 text-base pt-8 text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    {@render children()}
</div>