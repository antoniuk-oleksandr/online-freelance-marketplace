<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="text-base flex items-center mb-2 text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
    {@render children()}
</div>