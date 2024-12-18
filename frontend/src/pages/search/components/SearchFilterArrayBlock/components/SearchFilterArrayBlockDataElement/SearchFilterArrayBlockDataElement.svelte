<script lang="ts">
    import type {Language} from "@/types/Language.ts";
    import type {Skill} from "@/types/Skill.ts";
    import type {Category} from "@/types/Category.ts";
    import SearchFilterArrayBlockDataElementLayout
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataElement/SearchFilterArrayBlockDataElementLayout.svelte";
    import SearchFilterArrayBlockDataElementText
        from "@/pages/search/components/SearchFilterArrayBlock/components/SearchFilterArrayBlockDataElementText/SearchFilterArrayBlockDataElementText.svelte";
    import type {SearchFilterArrayAttribute} from "@/types/SearchFilterArrayAttribute.ts";
    import Checkbox from "@/common-components/Checkbox/Checkbox.svelte";

    type SearchFilterArrayBlockDataElementProps = {
        data: Language | Skill | Category,
        selected: boolean,
        attribute: SearchFilterArrayAttribute,

    }

    const {data, selected, attribute}: SearchFilterArrayBlockDataElementProps = $props();
</script>

<SearchFilterArrayBlockDataElementLayout attribute={attribute} id={data.id}>
    <Checkbox selected={selected}/>
    <SearchFilterArrayBlockDataElementText text={data.name}/>
</SearchFilterArrayBlockDataElementLayout>