<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type PaperElementProps = LayoutProps & {
        styles?: string,
    }

    const {children, styles}: PaperElementProps = $props();
</script>

<div
        class={`${styles} p-6 rounded-lg ring-1 bg-light-palette-background-block dark:bg-dark-palette-background-block ring-light-palette-divider dark:ring-dark-palette-divider`}
>
    {@render children()}
</div>