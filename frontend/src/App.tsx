import "./App.css";
import { Button } from "./components/ui/button";
import { NavigationMenuComponent } from "./components/ui/custom/NavigationMenu";
// import { LoginForm } from "./components/ui/custom/Login";

// import { Page } from "./components/ui/custom/Page";
// import { AddProject } from "./components/ui/custom/Project";

function App() {
  return (
    <>
      <NavigationMenuComponent />
      <section className="flex flex-col min-h-[80vh] items-center justify-center py-20 bg-gray-50 dark:bg-gray-900">
        <h1 className="text-5xl font-extrabold tracking-tight text-gray-900 dark:text-white sm:text-5xl mb-4 max-w-[50%] text-center">
          The best way to manage your projects in the new age
        </h1>
        <p className="max-w-xl text-2xl text-gray-600 dark:text-gray-300 mb-8 text-center">
          Effortlessly manage your projects, track progress, and collaborate
          with your team. Get started today!
        </p>
        <Button size="lg" className="px-12 py-6 text-xl">
          Get Started
        </Button>
      </section>
    </>
  );
}

export default App;
