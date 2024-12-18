<script lang="ts">
    import type {LayoutProps} from "@/types/LayoutProps.ts";
    import {handleSearchFilterArrayBlockDataElementClick} from "@/pages/search/handlers.ts";
    import type {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";

    type SearchFilterArrayBlockDataElementLayoutProps = LayoutProps & {
        id: number,
        attribute: SearchFilterArrayAttribute,
    }

    const {children, id, attribute}: SearchFilterArrayBlockDataElementLayoutProps = $props();
</script>

<button
        onclick={() =>
            handleSearchFilterArrayBlockDataElementClick(id.toString(), attribute)
        }
        class="flex gap-x-3 items-center py-2 px-6">
    {@render children()}
</button>