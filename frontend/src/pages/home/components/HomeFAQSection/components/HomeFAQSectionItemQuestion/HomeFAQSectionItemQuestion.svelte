<script lang="ts">
    import HomeFAQSectionItemQuestionLayout
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItemQuestion/HomeFAQSectionItemQuestionLayout.svelte";
    import HomeFAQSectionItemQuestionText
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItemQuestionText/HomeFAQSectionItemQuestionText.svelte";
    import HomeFAQSectionItemQuestionIcon
        from "@/pages/home/components/HomeFAQSection/components/HomeFAQSectionItemQuestionIcon/HomeFAQSectionItemQuestionIcon.svelte";

    type HomeFAQSectionItemQuestionProps = {
        question: string;
    };

    const props: HomeFAQSectionItemQuestionProps = $props();
</script>

<HomeFAQSectionItemQuestionLayout>
    <HomeFAQSectionItemQuestionIcon/>
    <HomeFAQSectionItemQuestionText {...props}/>
</HomeFAQSectionItemQuestionLayout>