CXX=g++
CFLAGS= -static
INCLUDE= -I../tsf4g_lib/include
LIBS= -L../tsf4g_lib/lib/ -ltsf4g -lscew -lexpat -lpthread

BINARY = $(patsubst ../../%.cpp,%.o,$(wildcard ../../*.cpp))
TARGET = tgamelog

all:$(TARGET)

$(TARGET):$(BINARY)
	$(CXX) $? -o $@ $(INCLUDE) $(CFLAGS) $(LIBS)
%.o:../../%.cpp
	$(CXX) -c $< -o $@ $(INCLUDE) $(CFLAGS)

clean:
	@rm -rf ${TARGET} $(BINARY)
