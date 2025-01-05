<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type HomeSectionLayoutProps = LayoutProps & {
        fullHeight?: boolean,
    }

    const {children, fullHeight}: HomeSectionLayoutProps = $props();
</script>

<section
        class="{fullHeight ? 'h-fit' : 'h-fit xl:h-128 '} py-6 sm:p-6 sm:dark:bg-dark-palette-background-block sm:bg-light-palette-background-block rounded-xl flex items-center justify-center">
    <div class="text-center w-full flex flex-col items-center justify-center gap-y-12">
        {@render children()}
    </div>
</section>
