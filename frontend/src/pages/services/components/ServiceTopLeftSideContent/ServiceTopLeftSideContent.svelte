<script lang="ts">
    import ServiceTopLeftSideContentLayout
        from "@/pages/services/components/ServiceTopLeftSideContent/ServiceTopLeftSideContentLayout.svelte";
    import type {Service} from "@/types/Service.ts";
    import UserInfoBlock from "@/pages/users/components/UserInfoBlock/UserInfoBlock.svelte";
    import Carousel from "@/common-components/Carousel/Carousel.svelte";
    import DividerElement from "@/common-components/DividerElement/DividerElement.svelte";

    const {title, freelancer, images}: Service = $props();
</script>

<ServiceTopLeftSideContentLayout>
    <h2 class="text-2xl font-semibold mb-4">{title}</h2>
    <UserInfoBlock
            useLink={true}
            showLanguages={false}
            size={"small"}
            user={freelancer}
    />
    <DividerElement className={"my-8"}/>
    <Carousel slides={images}/>
</ServiceTopLeftSideContentLayout>