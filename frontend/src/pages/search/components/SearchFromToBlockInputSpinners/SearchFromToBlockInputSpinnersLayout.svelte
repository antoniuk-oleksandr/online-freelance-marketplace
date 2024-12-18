<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="absolute top-1 right-2 text-light-palette-text-secondary dark:text-dark-palette-text-secondary flex flex-col">
    {@render children()}
</div>