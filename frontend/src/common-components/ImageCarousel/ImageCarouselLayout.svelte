<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="relative w-full max-w-4xl mx-auto overflow-hidden">
    {@render children()}
</div>