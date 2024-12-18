<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<div class="p-6 w-full flex items-end flex-grow">
    {@render children()}
</div>