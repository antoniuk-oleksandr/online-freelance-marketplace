<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="grid grid-cols-1 lg:grid-cols-service gap-8 mt-8">
    {@render children()}
</div>