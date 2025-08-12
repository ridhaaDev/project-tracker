import "./App.css";
import { Button } from "@/components/ui/button";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
} from "@/components/ui/navigation-menu";
import { LoginForm } from "./components/ui/custom/Login";

function App() {
  return (
    <>
      <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
        <LoginForm />
      </div>
    </>
  );
}

export default App;
