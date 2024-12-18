<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="h-16 border-b border-light-palette-divider dark:border-dark-palette-divider items-center px-6 gap-x-3 font-semibold hidden lg:flex">
    {@render children()}
</div>