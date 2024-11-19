<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="min-h-app w-full flex flex-col font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary select-none justify-center items-center">
    {@render children()}
</div>