<script>
    import HomeSectionLayout from "@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte";
    import HomeSectionHeader from "@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte";
    import HomeSectionItemList from "@/pages/home/common-components/HomeSectionItemList/HomeSectionItemList.svelte";

    const steps = [
        {
            icon: "hugeicons:search-02",
            title: "Search for a Freelancer",
            description: "Use our search page to find skilled freelancers for your needs."
        },
        {
            icon: "hugeicons:credit-card",
            title: "Place an Order",
            description: "Pay securely and place your order-request to get started."
        },
        {
            icon: "hugeicons:bubble-chat",
            title: "Chat with the Freelancer",
            description: "Discuss project details and stay updated throughout."
        },
        {
            icon: "hugeicons:checkmark-circle-01",
            title: "Review & Accept Work",
            description: "Approve the completed work or order-request changes. Cancel if needed."
        }
    ];
</script>

<HomeSectionLayout>
    <HomeSectionHeader
            title="How It Works"
            description="Simple steps to get your work done"
    />
    <HomeSectionItemList arr={steps}/>
</HomeSectionLayout>