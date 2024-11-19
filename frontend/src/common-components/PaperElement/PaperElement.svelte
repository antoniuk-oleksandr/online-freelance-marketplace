<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type PaperElementProps = LayoutProps & {
        styles?: string,
    }

    const {children, styles}: PaperElementProps = $props();
</script>

<div
        class={`${styles} p-6 rounded-lg border border-light-palette-divider dark:border-dark-palette-divider`}
>
    {@render children()}
</div>