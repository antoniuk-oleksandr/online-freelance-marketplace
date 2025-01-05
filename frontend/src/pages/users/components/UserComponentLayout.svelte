<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>


<div class="flex flex-col gap-y-3 h-fit">
    {@render children()}
</div>