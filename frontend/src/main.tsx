import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router";
import "./index.css";
import App from "./App.tsx";
import { LoginForm } from "./components/ui/custom/Login.tsx";
import { AddProject } from "./components/ui/custom/CreateProject.tsx";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ViewProjects } from "./components/ui/custom/ViewProjects.tsx";
import { ViewProject } from "./components/ui/custom/ViewProject.tsx";
import { SignUpForm } from "./components/ui/custom/SignUp.tsx";

const queryClient = new QueryClient();

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<App />} />
           <Route path="/sign-up" element={<SignUpForm />} />
          <Route path="/login" element={<LoginForm />} />
          <Route path="/projects" element={<AddProject />} />
          <Route path="/view/projects" element={<ViewProjects />} />
          <Route path="/view/projects/:id" element={<ViewProject />} />
          {/* <Route path="register" element={<Register />} /> */}
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  </StrictMode>
);
