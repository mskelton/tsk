package test_utils

type Fixtures struct {
	db string
}

// /// Builds a command to run the `tsk` binary. This auto populates the
// // /// `DATABASE_URL` environment variable with the path to the test database.
// // pub fn build_cmd(&self) -> Command {
// //     let mut cmd = Command::new(get_cargo_bin("tsk"));
// //     cmd.env("DATABASE_URL", &self.db_path);
// //
// //     if self.colors {
// //         cmd.env("CLICOLOR_FORCE", "1").env("COLORTERM", "truecolor");
// //     }
// //
// //     cmd
// // }
// //
// // pub fn cmd(&self, args: &str) -> Command {
// //     let mut cmd = self.build_cmd();
// //     cmd.args(args.split_whitespace());
// //     cmd
// // }
// //
// // pub fn run(&self, args: &str) {
// //     self.cmd(args).output().expect("failed to run command");
// // }
// //
// // pub fn colors(mut self) -> Self {
// //     self.colors = true;
// //     self
// // }
//
// func NewFixtures(args string) string {
//
// }
//
// func (f *Fixtures) Cmd(args string) string {
// 	// cmd := f.buildCmd()
// 	// cmd.Args(args.split(" "))
// 	// cmd
// }
//
// func (f *Fixtures) Colors(args string) {
// 	// f.Cmd(args).Output()
// 	cmd := exec.Command("tsk", args)
// 	cmd.Environ
// }
