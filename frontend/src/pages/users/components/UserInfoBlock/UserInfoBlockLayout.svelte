<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    let {children}: LayoutProps = $props();
</script>

<div class="flex gap-3 items-center flex-col sm:flex-row">
    {@render children()}
</div>