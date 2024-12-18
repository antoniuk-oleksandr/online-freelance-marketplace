<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} :LayoutProps = $props();
</script>

<div class="p-6 w-full flex flex-grow place-items-end gap-x-6">
    {@render children()}
</div>