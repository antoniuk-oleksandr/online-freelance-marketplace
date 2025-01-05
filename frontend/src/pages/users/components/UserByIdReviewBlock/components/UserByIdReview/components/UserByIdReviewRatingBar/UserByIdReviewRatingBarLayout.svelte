<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children}: LayoutProps = $props();
</script>

<div class={"flex gap-x-3 items-center"}>
    {@render children()}
</div>