<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="flex font-semibold gap-x-2 items-center justify-center w-[167.4px] px-4 py-2 hover:bg-light-palette-action-hover hover:dark:bg-dark-palette-action-hover duration-200 ease-out">
    {@render children()}
</div>