<script>
    import {Router, Route} from "svelte-routing";
    import UserPage from "./pages/users/UserPage.svelte";
    import Home from "./pages/home/HomePage.svelte";
    import ServicePage from "@/pages/services/ServicePage.svelte";
    import AppLayout from "@/AppLayout.svelte";
    import SignInPage from "@/pages/sign-in/SignInPage.svelte";
    import NotFound from "@/common-components/NotFound/NotFound.svelte";
    import SignUpPage from "@/pages/sign-up/SignUpPage.svelte";
    import ConfirmEmailPage from "@/pages/confirm-email/ConfirmEmailPage.svelte";
    import ForgotPasswordPage from "@/pages/forgot-password/ForgotPasswordPage.svelte";
    import ResetPasswordPage from "@/pages/reset-password/ResetPasswordPage.svelte";
</script>

<Router>
    <AppLayout>
        <Route path="/">
            <Home/>
        </Route>
        <Route path="/services/:id" let:params>
            <ServicePage id={params.id}/>
        </Route>
        <Route path="/users/:id" let:params>
            <UserPage id={params.id}/>
        </Route>
        <Route path="/sign-in">
            <SignInPage/>
        </Route>
        <Route path="/sign-up">
            <SignUpPage/>
        </Route>
        <Route path="/confirm-email">
            <ConfirmEmailPage/>
        </Route>
        <Route path="/forgot-password">
            <ForgotPasswordPage/>
        </Route>
        <Route path="/reset-password">
            <ResetPasswordPage/>
        </Route>
        <Route path="/*">
            <NotFound/>
        </Route>
    </AppLayout>
</Router>
