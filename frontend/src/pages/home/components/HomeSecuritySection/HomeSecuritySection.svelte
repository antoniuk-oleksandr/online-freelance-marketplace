<script>
    import HomeSectionLayout from "@/pages/home/common-components/HomeSectionLayout/HomeSectionLayout.svelte";
    import HomeSectionHeader from "@/pages/home/common-components/HomeSectionHeader/HomeSectionHeader.svelte";
    import HomeSectionItemList from "@/pages/home/common-components/HomeSectionItemList/HomeSectionItemList.svelte";

    const trustFeatures = [
        {
            icon: "hugeicons:blocked",
            title: "Secure Payments",
            description: "All transactions are encrypted and handled securely to protect your financial information."
        },
        {
            icon: "hugeicons:user-check-01",
            title: "Verified Freelancers",
            description: "Our freelancers go through a rigorous verification process to ensure quality and reliability."
        },
        {
            icon: "hugeicons:clock-02",
            title: "Timely Deliveries",
            description: "Our platform ensures deadlines are respected and your projects are delivered on time."
        },
        {
            icon: "hugeicons:bubble-chat-done",
            title: "24/7 Support",
            description: "Our dedicated support team is available around the clock to assist you with any queries."
        }
    ];
</script>

<HomeSectionLayout>
    <HomeSectionHeader
        title="Trust & Security"
        description="We prioritize your trust and safety while ensuring the best experience on our platform."
        />
    <HomeSectionItemList arr={trustFeatures}/>
</HomeSectionLayout>