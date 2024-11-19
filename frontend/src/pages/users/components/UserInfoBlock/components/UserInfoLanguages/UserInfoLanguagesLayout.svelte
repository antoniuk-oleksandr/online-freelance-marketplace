<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children} : LayoutProps = $props()
</script>

<div class="flex items-center text-lg text-light-palette-text-secondary dark:text-dark-palette-text-secondary gap-x-2">
    {@render children()}
</div>