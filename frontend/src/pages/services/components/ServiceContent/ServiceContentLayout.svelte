<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class="grid grid-cols-1 lg:grid-cols-service gap-6 mt-6">
    {@render children()}
</div>