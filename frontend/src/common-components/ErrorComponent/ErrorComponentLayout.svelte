<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="fixed bg-light-palette-background-block dark:bg-dark-palette-background-block h-svh w-svw grid place-items-center z-[100]">
    {@render children()}
</div>