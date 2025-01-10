<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type CardInputElementProps = LayoutProps & {
        styles?: string,
    }

    const {children, styles}: CardInputElementProps = $props();
</script>

<div class="{styles} w-full font-medium flex flex-col gap-1">
    {@render children()}
</div>