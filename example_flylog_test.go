package flylog_test

import "github.com/flylog/flylog"

func Example() {
	logs := flylog.GetLogger()
	for i := 0; i < 100; i++ {
		logs.Debug("example flylog test", i)
		logs.Info("example lylog test2", i)
		logs.Error("example lylog test2", i)
		logs.Warn("example lylog test2", i)
	}

	// Output:
	// [2021-08-08 08:08:08] -- main.main:12 -- example lylog test295 [Info] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:13 -- example lylog test295 [Error] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:14 -- example lylog test295 [Warn] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:11 -- example flylog test96 [Debug] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:12 -- example lylog test296 [Info] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:13 -- example lylog test296 [Error] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:14 -- example lylog test296 [Warn] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:11 -- example flylog test97 [Debug] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:12 -- example lylog test297 [Info] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:13 -- example lylog test297 [Error] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:14 -- example lylog test297 [Warn] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:11 -- example flylog test98 [Debug] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:12 -- example lylog test298 [Info] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:13 -- example lylog test298 [Error] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:14 -- example lylog test298 [Warn] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:11 -- example flylog test99 [Debug] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:12 -- example lylog test299 [Info] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:13 -- example lylog test299 [Error] <FLYLOG>
	// [2021-08-08 08:08:08] -- main.main:14 -- example lylog test299 [Warn] <FLYLOG>
}
