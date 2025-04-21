<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type ModalHeaderLayoutProps = LayoutProps & {
        headerStyles?: string,
    }

    const {children, headerStyles}: ModalHeaderLayoutProps = $props();
</script>

<div class="{headerStyles} p-6 text-lg font-semibold border-b border-light-palette-divider dark:border-dark-palette-divider flex justify-between items-center">
    {@render children()}
</div>
