<script lang="ts">
    import {handleSearchFilterArrayBlockClick} from "@/pages/search/handlers.js";
    import type {Language} from "@/types/Language.ts";
    import type {Skill} from "@/types/Skill.ts";
    import type {Category} from "@/types/Category.ts";
    import type {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";
    import type {LayoutProps} from "@/types/LayoutProps.ts";

    type SearchFilterArrayBlockLayoutProps = LayoutProps & {
        data: Language[] | Skill[] | Category[],
        title: string,
        searchParamsAttribute: SearchFilterArrayAttribute,
    }

    const {title, data, searchParamsAttribute, children}: SearchFilterArrayBlockLayoutProps = $props();

</script>

<button
        onclick={() => handleSearchFilterArrayBlockClick(data, title, searchParamsAttribute)}
        class="flex font-semibold justify-between items-center py-2 px-6 cursor-pointer duration-200 ease-out rounded-l  hover:bg-light-palette-action-hover hover:dark:bg-dark-palette-action-hover">
    {@render children()}
</button>