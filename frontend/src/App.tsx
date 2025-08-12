import "./App.css";
import { Button } from "@/components/ui/button";
import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuList,
} from "@/components/ui/navigation-menu";

function App() {
  return (
    <>
      <div>
        <NavigationMenu>
          <NavigationMenuList>
            <NavigationMenuItem>
              Home
            </NavigationMenuItem>

             <NavigationMenuItem>
              Projects
            </NavigationMenuItem>

            <NavigationMenuItem>
              Sign in
            </NavigationMenuItem>



          </NavigationMenuList>
        </NavigationMenu>
        <Button>Click me</Button>
      </div>
    </>
  );
}

export default App;
