<script lang="ts">
    import Icon from "@iconify/svelte";
    import type {Language} from "@/types/Language.ts";
    import type {Skill} from "@/types/Skill.ts";
    import type {Category} from "@/types/Category.ts";
    import type {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";
    import SearchFilterArrayBlockLayout
        from "@/pages/search/components/SearchFilterArrayBlock/SearchFilterArrayBlockLayout.svelte";
    import type {SearchPageParams} from "@/types/SearchPageParams.ts";
    import SearchFilterArrayBlockSelectedDisplay
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockSelectedDisplay/SearchFilterArrayBlockSelectedDisplay.svelte";

    type SearchFilterArrayBlockProps = {
        data: Language[] | Skill[] | Category[],
        title: string,
        searchParamsAttribute: SearchFilterArrayAttribute,
        searchPageParams: SearchPageParams,
    }

    const props: SearchFilterArrayBlockProps = $props();
</script>

<SearchFilterArrayBlockLayout {...props}>
    <span>{props.title}</span>
    <div class="text-light-palette-text-secondary dark:text-dark-palette-text-secondary flex gap-x-1 items-center">
        <SearchFilterArrayBlockSelectedDisplay {...props}/>
        <Icon
                icon="fe:arrow-right"
                width={18}
                height={18}
        />
    </div>
</SearchFilterArrayBlockLayout>