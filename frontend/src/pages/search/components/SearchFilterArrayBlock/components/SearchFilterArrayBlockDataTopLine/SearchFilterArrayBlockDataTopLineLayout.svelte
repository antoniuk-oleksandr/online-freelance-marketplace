<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex font-semibold px-6 py-3 items-center justify-between border-b h-16 border-light-palette-divider dark:border-dark-palette-divider">
    {@render children()}
</div>