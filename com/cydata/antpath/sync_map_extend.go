/**
 * Created by GoLand.
 * Brief: matcher fields read/write
 * User: vibrant
 * Date: 2019/04/03
 * Time: 11:20
 */
package antpath

//ClearSyncMap
func ClearSyncMap(m *SyncMap) {
	if m != nil {
		m.MyRange(func(key, value interface{}) bool {
			m.MyDelete(key)
			return true
		})
	}
}
