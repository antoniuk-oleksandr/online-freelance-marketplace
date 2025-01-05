<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import PaperElement from "@/common-components/PaperElement/PaperElement.svelte";
    import {Link} from "svelte-routing";
    import {getHoverClass} from "@/utils/utils.ts";

    type CategoriesSectionItemLayoutProps = LayoutProps & {
        id: number
    }

    const {children, id}: CategoriesSectionItemLayoutProps = $props();

    const hoverClass = getHoverClass();
</script>

<Link to="/search?category={id}">
    <PaperElement
            styles="aspect-square !p-3 grid-item flex-none {hoverClass} hover:bg-light-palette-hover flex flex-col items-center justify-center gap-2 h-32 sm:h-36">
        {@render children()}
    </PaperElement>
</Link>
