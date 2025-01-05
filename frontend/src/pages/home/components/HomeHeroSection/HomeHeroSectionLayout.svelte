<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    const {children} : LayoutProps = $props();
</script>

<section class="relative h-128 rounded-2xl overflow-hidden bg-gradient-to-r from-cyan-500 to-blue-500 w-full text-white flex flex-col justify-center items-center px-6">
    {@render children()}
</section>