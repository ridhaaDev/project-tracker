import "./App.css";
// import { LoginForm } from "./components/ui/custom/Login";

import { Page } from "./components/ui/custom/Page";
import { AddProject } from "./components/ui/custom/Project";

function App() {
  return (
    <>
      <Page>
        {/* <LoginForm /> */}

        <AddProject />
      </Page>
    </>
  );
}

export default App;
