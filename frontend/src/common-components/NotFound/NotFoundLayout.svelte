<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="absolute left-0 top-0 h-svh w-full bg-light-palette-background-default dark:bg-dark-palette-background-default flex flex-col font-semibold text-light-palette-text-primary dark:text-dark-palette-text-primary select-none justify-center items-center">
    {@render children()}
</div>