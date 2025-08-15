import { Link, useLocation } from "react-router";
import {
  NavigationMenu,
  NavigationMenuContent,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  NavigationMenuTrigger,
} from "../navigation-menu";
import { Button } from "../button";

export const NavigationMenuComponent = () => {
  const location = useLocation();

  return (
    <div className="flex justify-between h-20 p-4 bg-white drop-shadow-md">
      <NavigationMenu viewport={false}>
        <NavigationMenuList>
          {location.pathname == "/login" ? (
            <></>
          ) : (
            <>
              {" "}
              <NavigationMenuItem>
                <NavigationMenuLink asChild>
                  <Link to={{ pathname: "/" }}>Home</Link>
                </NavigationMenuLink>
              </NavigationMenuItem>
              <NavigationMenuItem>
                <NavigationMenuTrigger>Project</NavigationMenuTrigger>
                <NavigationMenuContent>
                  <ul className="grid w-[200px] gap-4">
                    <li>
                      <NavigationMenuLink asChild>
                        <Link
                          to={{
                            pathname: "/projects",
                          }}
                        >
                          Create new Project
                        </Link>
                      </NavigationMenuLink>
                      <NavigationMenuLink asChild>
                        <Link
                          to={{
                            pathname: "/docs",
                          }}
                        >
                          Documentation
                        </Link>
                      </NavigationMenuLink>
                      <NavigationMenuLink asChild>
                        <Link
                          to={{
                            pathname: "/docs",
                          }}
                        >
                          Blocks
                        </Link>
                      </NavigationMenuLink>
                    </li>
                  </ul>
                </NavigationMenuContent>
              </NavigationMenuItem>
            </>
          )}
        </NavigationMenuList>
      </NavigationMenu>

      <NavigationMenu>
        <NavigationMenuList>
          {location.pathname == "/login" ? (
            <>
             <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Button>
                  <Link
                    to={{
                      pathname: "/",
                    }}
                  >
                    Home
                  </Link>
                </Button>
              </NavigationMenuLink>
            </NavigationMenuItem>
            </>
          ) : (
            <NavigationMenuItem>
              <NavigationMenuLink asChild>
                <Button>
                  <Link
                    to={{
                      pathname: "/login",
                    }}
                  >
                    Login
                  </Link>
                </Button>
              </NavigationMenuLink>
            </NavigationMenuItem>
          )}
        </NavigationMenuList>
      </NavigationMenu>
    </div>
  );
};
