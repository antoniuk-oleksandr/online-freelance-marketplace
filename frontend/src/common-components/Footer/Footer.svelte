<script>
    import FooterLayout from "@/common-components/Footer/FooterLayout.svelte";
    import AppLogo from "@/common-components/AppLogo/AppLogo.svelte";
    import FooterIcon from "@/common-components/Footer/components/FooterIcon/FooterIcon.svelte";
</script>

<FooterLayout>
    <div class={"flex gap-x-6 justify-center md:justify-start items-center w-full flex-col sm:flex-row gap-y-4"}>
        <AppLogo/>
        <p>© {(new Date()).getUTCFullYear()} Online Freelance Marketplace. All rights reserved.</p>
    </div>
    <div class={"text-xl flex gap-x-6"}>
        <FooterIcon iconName={"hugeicons:instagram"} link={"https://www.instagram.com/"}/>
        <FooterIcon iconName={"hugeicons:facebook-01"} link={"https://www.facebook.com/"}/>
        <FooterIcon iconName={"hugeicons:linkedin-01"} link={"https://www.linkedin.com/"}/>
        <FooterIcon iconName={"hugeicons:new-twitter-rectangle"} link={"https://x.com/"}/>
    </div>
</FooterLayout>
