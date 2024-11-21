<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="size-full h-svh grid place-items-center">
    {@render children()}
</div>