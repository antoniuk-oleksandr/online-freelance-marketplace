<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="flex flex-col gap-y-6 w-full">
    {@render children()}
</div>