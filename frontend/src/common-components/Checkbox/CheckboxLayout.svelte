<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="rounded-sm text-cyan-500 grid place-items-center size-4.5 ring-1 ring-light-palette-divider dark:ring-dark-palette-divider duration-200 ease-out hover:bg-light-palette-action-hover hover:dark:bg-dark-palette-action-hover">
    {@render children()}
</div>